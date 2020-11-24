set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${CMAKE_CURRENT_SOURCE_DIR}/cmake")

cmake_minimum_required(VERSION 3.0)

set(project_version "1.0.1")

project(redtun Go)

include(cmake/golang.cmake)
#include(cmake/flags.cmake) if use CGO

set(output_dir ${CMAKE_CURRENT_BINARY_DIR}/bin)

# add compile source dir
add_subdirectory(cmd/RedTun-Proxy ${output_dir})

set (default_config_uri ${output_dir})

#copy default config file to output
message("-- Move the default configuration file to the project compilation output directory.")

file(COPY ${CMAKE_CURRENT_LIST_DIR}/example/rd_config.yaml
     DESTINATION   ${default_config_uri})

#import rpm pack
set(package_type RPM)

include(cmake/cpack.cmake)
