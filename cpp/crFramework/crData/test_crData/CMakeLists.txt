# CMake 最低版本号要求
cmake_minimum_required (VERSION 3.6)

# 项目信息
project (test_crData)


include_directories(../include)

aux_source_directory(./ DIR_SRC)
aux_source_directory(../src CRDATA_SRC)

# 指定生成目标
add_executable(test_crData ${DIR_SRC} ${CRDATA_SRC})

set_property(TARGET test_crData PROPERTY CXX_STANDARD 11)