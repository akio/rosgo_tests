#!/usr/bin/env python
import rospy
from rosgo_test.msg import Hello


def callback(data):
    rospy.loginfo(rospy.get_name() + ": I heard %s" % data.data)


def listener():
    rospy.init_node('listener', anonymous=True, log_level=rospy.DEBUG)
    rospy.Subscriber("chatter", Hello, callback)
    rospy.spin()


if __name__ == '__main__':
    listener()
