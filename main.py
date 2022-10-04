import os
import time
import pathlib
from pathlib import Path


def mijn_sum(file_name):
	n = input("Введите число слагаемых n: ")
	m = input("Введите степень m: ")
	try:
		n = int(n)
		m = int(m)

		f = open('data.txt','w')
		f.write(f'{n}\n')
		f.write(f'{m}')
		
		f.close()

	except ValueError:
		print("Это не правильный ввод. Это не число вообще! Это строка, попробуйте еще раз.")

	os.startfile(file_name)
	# Этот метод работает так, как будто файл запустили двойным кликом -- 
	# программы будут запущены, для файлов будет вызвана ассоциированная 
	# с ними программа.

	time.sleep(0.1)

	with open('data.txt') as f:
		print(f"sum = {f.readline()}")

	file=pathlib.Path(f"{Path.cwd()}/data.txt") # удаляем файл из текущей директории
	file.unlink()


if __name__ == "__main__":
	mijn_sum("sumc.exe")
	mijn_sum("sumgo.exe")