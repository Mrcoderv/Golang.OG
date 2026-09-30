// building the cli in go .
the basic concept of building a cli in go is to use the "flag" package to parse command-line arguments and flags. You can define flags using the flag package, and then use them to control the behavior of your program.
eg if we are building the algoritm that is used to download the software so 
command line arguments can be used to specify the software name, version, and download location.
 The flag package allows you to define flags for each of these parameters, and then parse them from the command line when the program is run.
 // command line argument is a way to pass information to a program when it is executed from the command line. In Go, you can use the "flag" package to handle command line a  rguments.

2. os.Args  == raw actual agrs  that are slices 
os.Args
   ↓
┌──────────────┬────────┬──────┬───────┐
│ program name │ Raghav │  20  │ Nepal │
└──────────────┴────────┴──────┴───────┘
      [0]         [1]     [2]    [3]
it treat the value as the arrary so they are sepertated through the space 
so we can access the valuse through the sequense of index that are start from the 0 ,1 ,2  
whera fist 0 position is trated as the file name .

1.flag Package  == Named arguments

THIS IS THE PACKAHGE THAT  that hst the power of the standard library for handling command-line flags.


flag is used to give the name of the arugemnt like providing the categorical name to the value .
flag.string is used to create the name of the argument 
flag.Parse() is used to break the  the command-line flags provided by the user. It processes the flags defined using flag.String, flag.Int, and other flag functions, and assigns the values to the corresponding variables. After calling flag.Parse(), you can access the values of the flags through the pointers returned by the flag functions.
 frag.Sting give the pointer value so that * is use d to concvert the dat aon the value 
 flag uses the defult value 
 so the the data will not be null and the .

 Step 4 — Subcommands  
 like creating the self command for the system 
 ie. creating the admin of the system 
 
