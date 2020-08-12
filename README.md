# Download-Go

#### Introduction
Concurrent download program written by go

#### Instructions for use

|Command|Parameter|Description|
|----|----|----|
|-n|int|Set the number of threads (the default is the number of logical CPUs)|
|-t|int|Set the timeout time, in seconds, -1 means unlimited timeout (default -1)|
|-u|string|Set link (required)|
|-p|string|Set the setting save path (default is the current directory)|
|-f|string|Set save file name|
|-tx| |Start downloading from the Config.txt file in the current directory|
|-h| |Output command list|
|-v| |Output program information|