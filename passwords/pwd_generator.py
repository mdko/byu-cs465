import argparse
import crypt
import time


def generate_files():
	parser = argparse.ArgumentParser()
	parser.add_argument("file_name")
	args = parser.parse_args()
	file_name = args.file_name
	
	inputfile = open(file_name)
	pwdfile = open('passwd.txt', 'r+')
	shadowfile = open('shadow.txt', 'r+')

	for line in inputfile:
		username, password = line.split(" ")
		hashedpwd = crypt.crypt(password, crypt.METHOD_MD5)
		#		  username:pwd:uid:groupid:uidinfo:homedir:shell
		pwdline = username + ':x:1001:1001::test:directory:shell12\n'
		#		  username:pwd:lastchanged:minimum:maximum:warn:inactive:expire
		days = int(time.time() / 86400)
		shdwline = username + ':' + hashedpwd + ':' + str(days) + ':0:99999:7:::\n'

		pwdfile.write(pwdline)
		shadowfile.write(shdwline)

	pwdfile.close()
	shadowfile.close()

def run():
	generate_files()


if __name__ == "__main__":
	run()