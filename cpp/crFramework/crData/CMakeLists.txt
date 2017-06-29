
cmake_minimum_required (VERSION 3.6)

project(crData)

include_directories(./include)
aux_source_directory(./src DIR_SRC)

set(LIBRARY_OUTPUT_PATH ../../bin)

add_library(crData SHARED ${DIR_SRC})