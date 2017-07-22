cmake_minimum_required(VERSION 3.7)
project(leetcode)

set(CMAKE_CXX_STANDARD 11)

add_subdirectory(lib/googletest)
include_directories(lib/googletest/googletest/include)
include_directories(lib/googletest/googlemock/include)

file(GLOB_RECURSE HEADER_ALG "algorithms/*.hpp")

set(HERDER_FILES ${HEADER_ALG})

file(GLOB_RECURSE SOURCE_ALG "algorithms/*.cpp")

set(SOURCE_FILES main_test.cpp ${SOURCE_ALG})

add_executable(main_test ${SOURCE_FILES} ${HERDER_FILES})
target_link_libraries(main_test gtest gtest_main)
target_compile_options(main_test PRIVATE -fno-stack-protector)