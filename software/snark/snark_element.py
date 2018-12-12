import sys, os

if sys.argv[1]=="setup":
  os.system("./zokrates compile -i snark_element.py > /dev/null")
  os.system("./zokrates setup > /dev/null")
  os.system("./zokrates export-verifier > /dev/null")
  f=open("verifier.sol")
  data=f.read()
  f.close()
  data=data.replace(") public returns (bool) {", ") public view returns (bool) {")
  f=open("verifier.sol", "w")
  f.write(data)
  f.close()

if sys.argv[1]=="prove":
  os.system("./zokrates compute-witness -a `cat args.txt`")
  os.system("./zokrates export")



