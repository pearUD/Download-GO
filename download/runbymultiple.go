package download

import (
	"download-go/entity"
	"download-go/response"
	"io"
	"log"
	"strconv"
	"sync"
)

var (
	m   chan entity.Item
	w   chan *entity.WItem
	wg  sync.WaitGroup
	res entity.Resources
)

func RunByMultiple(r entity.Resources, conf entity.Conf) {
	r.SetContent(r.GetResponse().ContentLength)
	m = make(chan entity.Item, conf.GetCpu())
	w = make(chan *entity.WItem, 10)
	res = r
	itemSize := r.GetContent() / int64(conf.GetCpu())
	if r.GetContent()%int64(conf.GetCpu()) > 0 {
		itemSize++
	}
	var i, start, end int64
	go setItem()
	for i = 0; i < int64(conf.GetCpu()); i++ {
		request := response.CreateRequest(r.GetUrl(), "GET", nil)
		start = i * itemSize
		end = start + itemSize
		if end > r.GetContent() {
			end = r.GetContent()
		}
		response.ChangeHead(request, "Range", "bytes="+strconv.FormatInt(start, 10)+"-"+strconv.FormatInt(end-1, 10))
		wg.Add(1)
		m <- entity.Item{Request: request, Start: start, End: end, File: r.GetFile()}
		go loadItem()
	}
	wg.Wait()
	close(w)
	wg.Add(1)
	wg.Wait()
}

func loadItem() {
	item := <-m
	resp := response.GetResponse(item.Request)
	start := item.Start
	end := item.End
	buf := [buffer * 1024]byte{}
	for {
		rs, err := resp.Body.Read(buf[:])
		if rs > 0 {
			w <- &entity.WItem{Start: start, Length: rs, Buffer: buf}
			start += int64(rs)
			if start > end {
				log.Printf("写入错误%d %d", resp.ContentLength, end-start)
			}
		}
		if err != nil {
			if err != io.EOF {
				log.Fatalf("文件读取时发时错误 Error：%s", err.Error())
			}
			break
		}
	}
	defer func() {
		wg.Done()
		resp.Body.Close()
	}()
}

func setItem() {
	for {
		wItem, ok := <-w
		if wItem != nil {
			ws, err := res.GetFile().WriteAt(wItem.Buffer[:wItem.Length], wItem.Start)
			if err != nil {
				log.Fatalf("文件写入时出现错误 Error：%s", err.Error())
			}
			if ws != wItem.Length {
				log.Fatalf("文件写入时出现错误 Error：读取与写入文件不一致")
			}
			res.SetInLoad(res.GetInLoad() + int64(ws))
		}
		if !ok && len(w) == 0 {
			break
		}
	}
	defer func() {
		_ = res.GetFile().Close()
		wg.Done()
	}()
}
