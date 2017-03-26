#!/usr/bin/python3
# Copyright (C) 2017 Paul Kramme

show_success = True

import zlib
import sys
import os

class color:
	HEADER = '\033[95m'
	OKBLUE = '\033[94m'
	OKGREEN = '\033[92m'
	WARNING = '\033[93m'
	FAIL = '\033[91m'
	ENDC = '\033[0m'
	BOLD = '\033[1m'
	UNDERLINE = '\033[4m'


def split(string, splitters):
	final = [string]
	for x in splitters:
		for i,s in enumerate(final):
			if x in s and x != s:
				left, right = s.split(x, 1)
				final[i] = left
				final.insert(i + 1, x)
				final.insert(i + 2, right)
	return final


def crc(filepath):
	previous = 0
	try:
		for line in open(filepath,"rb"):
			previous = zlib.crc32(line, previous)
	except OSError:
		print("CRC ERROR: OSERROR")
	return "%X"%(previous & 0xFFFFFFFF)


def scandirectory(walk_dir, scanfile, verbose = False):
	try:
		current_scan = []
		for root, subdirs, files in os.walk(walk_dir):
			#current_scan.extend([f"{root}\n"])
			for filename in files:
				if filename == "media-validation-tool" or filename == "media.csv" or filename == "media-validation-tool.py" or filename == "." or filename == "..":
					break
				else:
					if os.name == "posix":
						file_path = root + "/" + filename
					elif os.name == "nt":
						file_path = root + "\\" + filename
					else:
						pass
					checksum = crc(os.path.join(root, filename))
					current_scan.extend([file_path + "," + checksum + ",\n"])
		with open(scanfile, "w") as current_scan_file:
			current_scan_file.writelines(current_scan)			
	except FileNotFoundError:
		if verbose == True:
			print("SCAN ERROR: FILE NOT FOUND")

def validate(path = ".", verifile = "./media.csv"):
	print("Validating " + path)
	f = open(verifile, "r")
	verifileslist = f.readlines()
	scandirectory(".", "media.csv.1")
	fold = open("media.csv", "r")
	fnew = open("media.csv.1", "r")
	foldarray = fold.readlines()
	fnewarray = fnew.readlines()
	errorcount = 0
	successcount = 0
	for line in foldarray:
		splitline = split(line, ",")
		if not line in fnewarray:
			print("* ERROR " + splitline[0])
			errorcount = errorcount + 1
		else:
			if show_success == True:
				print("SUCCESS " + splitline[0])
			successcount = successcount + 1
	print(errorcount, "errors in", successcount + errorcount, "files")
	fold.close()
	fnew.close()
	os.remove("media.csv.1")

def create(path = ".", verifile = "./media.csv"):
	print("Creating media file...")
	scandirectory(".", "media.csv")
	
def main():
	print("MEDIA VALIDATION TOOL")
	print("(C) 2017 Paul Kramme")
	if len(sys.argv) < 2:
		validate()
		input("Press ENTER to continue...")
	else:
		create()
	
	

if __name__ == "__main__":
	main()
