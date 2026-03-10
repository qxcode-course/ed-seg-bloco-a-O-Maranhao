def fuga():
     h=int(input)
     p=int(input)
     f=int(input)
     d=int(input)

     while 1>0:
         f+=d
         if f== '-1':
             f=15
         if f=='16':
             f=0
         if f==h:
             print("S")
             break
         if f==p:
             print("N")
             break
