cmake_minimum_required(VERSION 2.8.3)
project(rosgo_tests)

find_package(
  catkin REQUIRED
  COMPONENTS
    message_generation
    rosgo_core
    rospy
    std_msgs
)


add_message_files(
  FILES
  Hello.msg
  AllFieldTypes.msg
)


add_service_files(
  FILES
  AddTwoInts.srv
)


generate_messages(DEPENDENCIES std_msgs)


catkin_package(
  CATKIN_DEPENDS message_runtime rospy rosgo_core
)


catkin_add_go_executable(
    ${PROJECT_NAME}/test_talker
    DEPENDS ${catkin_GO_LIBRARIES} ${PROJECT_NAME}_gengo)

catkin_add_go_executable(
    ${PROJECT_NAME}/test_talker_with_callbacks
    DEPENDS ${catkin_GO_LIBRARIES} ${PROJECT_NAME}_gengo)

catkin_add_go_executable(
    ${PROJECT_NAME}/test_listener
    DEPENDS ${catkin_GO_LIBRARIES} ${PROJECT_NAME}_gengo)

catkin_add_go_executable(
    ${PROJECT_NAME}/test_listener_with_event
    DEPENDS ${catkin_GO_LIBRARIES} ${PROJECT_NAME}_gengo)

catkin_add_go_executable(
    ${PROJECT_NAME}/test_client
    DEPENDS ${catkin_GO_LIBRARIES} ${PROJECT_NAME}_gengo)

catkin_add_go_executable(
    ${PROJECT_NAME}/test_server
    DEPENDS ${catkin_GO_LIBRARIES} ${PROJECT_NAME}_gengo)

catkin_add_go_test(
    ${PROJECT_NAME}/test_message
    DEPENDS ${catkin_GO_LIBRARIES} ${PROJECT_NAME}_gengo)
