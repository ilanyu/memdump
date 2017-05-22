
# memdump

example use: 

	# ps // get pid 3149
	# cat /proc/3149/maps // get b6946000-b6a12000 r-xp 00000000 08:06 652        /system/lib/libsqlite.so
	# ./memdump -pid 3149 -saddr b6946000 -eaddr b6a12000
