a=int(raw_input('Enter first number : '))
b=int(raw_input('Enter 2nd   number : '))
c=int(raw_input('Enter 3rd   number : '))

if a==b ||b==c||c==a:
	print "\nAll the integers should be distinct...then we can campare the integers...so please try again"

print "\n"

if a>b:
	#check a and c as it's sure a>b
	if a>c:
		print a," is greater";
	elif c>b:
		print c," is greater";
	else:
		print b," is greater";
elif b>c:
	print b," is greater";
elif a>c:
	print a," is greater";
else:
	print c," is greater";
