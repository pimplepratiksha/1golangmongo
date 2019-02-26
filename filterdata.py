if __name__=="__main__":
    fp1=open("r.json","r")
    fp2=open("newset.json","w+")
    fp=fp1.readline()
    while(fp):
        start=0
        index=fp.find("_id")
        arr=fp[start:index-1]
        fp2.write(arr)
        endline=fp.find('\n')
        arr=fp[index:endline]
        end=arr.find(",")
        arr=arr[end+1:endline]
        fp2.write(arr)
        fp2.write(",\n")
        fp=fp1.readline()
