#!/usr/bin/env python
import rospy
from rosgo_test.msg import Hello


def talker():
    pub = rospy.Publisher('chatter', Hello)
    rospy.init_node('talker', anonymous=True)
    while not rospy.is_shutdown():
        str = "%s: hello world %s" % (rospy.get_name(), rospy.get_time())
        rospy.loginfo(str)
        pub.publish(Hello(str))
        rospy.sleep(1.0)


if __name__ == '__main__':
    try:
        talker()
    except rospy.ROSInterruptException:
        pass
