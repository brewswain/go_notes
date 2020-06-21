# Custom Functions

Functions are great. I know this, you(I) know this, the world knows this. So let's create some of our own functions. Of course, these will be leveraging some of the interesting things about go that we've learnt, like how to return multiple values. We'll also go over how to declare functions both with and without parameters, and finally, we'll learn about `pointers`.

Pointers are brand new to me, but they're used to make more memory-efficient function calls.

Let's begin by discussing a problem where a function would be super useful--suppose we have a problem that'll need to run the exact same calculations but with different parameters. That's pretty cut and dry, so let's look at this example:

```
func main() {
    width = 4.2
    height = 3.0
    area = width * height
    fmt.Println(area/10.0, "liters needed.")

    width = 5.2
    height = 3.5
    area = width * height
    fmt.Println(area/10.0, "liters needed.")
}
```

Again, pretty obvious problem here. Something that may not be too obvious to you (depending on your knowledge of floats), is that we'll get an over-detailed _yet_ slightly innacurate response. This is something that's not limited to Go and honestly not that interesting, [so feel free to do some research on it yourself](http://effbot.org/pyfaq/why-are-floating-point-calculations-so-inaccurate.htm#:~:text=It's%20a%20problem%20caused%20by,resulting%20in%20small%20roundoff%20errors).

I know this isn't related to functions directly, but it's related to some useful formatting tips, which is super useful. If you want to just skip over to actually learning function declaration et al, look at the next file.

Anyway, the answer to this problem is usually to round them down. Let's look at the function that you may or may not already be using (I sure as hell do) for that!

## Formatting Output with Printf and Sprintf

Instead of using the innacurate method of:

```
fmt.Println("About one-third:",  1.0/3.0)


```

To format our printed output , we can use `Printf`. As you've no doubt gleaned from the name, `Printf` stands for "Print, with formatting":

```
fmt.Println("About one-third: %0.2f\n",  1.0/3.0)

```

Respectively, our two options would output something like:

```
About one-third: 0.3333333333333333
About one-third: 0.33
```

Obivously, our second option is desirable.

Now, onto `Sprintf`. This function works just like `Printf`, except that it returns a formatted string instead of printing it:

```
resultString := fmt.SPrintf("About one-third: %0.2f\n",  1.0/3.0)
fmt.Printf(resultString)
```

Again, we'd get 0.33 printed out. Obviously though, the syntax looks a little funky, so let's look at 2 of the features that we used--Formatting verbs, and value widths.

## Formatting Verbs

The first argument to `Printf` is a string that'll be used to format the output, most of which looks normal. However, any time we use `%`, it'll indicate the start of a formatting verb, which is a section of the string that'll be substituted with a value in a particular format:

```
fmt.Printf("the %s cost %d cents.\n" "gumballs", 23)
```

```
The gumballs will cost 23 cents each.
```

Now to the actual important part of all of this, the cheat sheet(Again, this is a fairly shallow look at all this since this is fairly common):

```
%f - Floating-point number
%d - Decimal integer
%s - String
%t - Boolean
%v - Any value(Chooses an appropriate format based on the supplied value's type)
%#v - Any value, formatted as it would appear in Go
%T - Type of the supplied value
%% - An actual percent sign
```

### Example cases

```
fmt.Printf("A float: %f\n", 3.1415)
fmt.Printf("An integer: %d\n" 15)
fmt.Printf("A string: %s\n", "hello")
fmt.Printf("A boolean: %t\n", true)
fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
fmt.Printf("Percent sign: %%\n")
```

Note that the `Values` section would output some different stuff!
The first use would output:

```
1.2      true
```

But the second one would output:

```
1.2 "\t" true
```

Also, you probably noted that we're manually adding a newline everytime. This is since unlike `Println`, `Printf` doesn't automatically add a newline for us.

## Formatting value widths

For situations when you want to round your value to a specific number, `formatting verbs` will allow us to specify the width of the formated value.

For instance, if we want to format some data in a plain-text table, we'd need to ensure that the formatted value fills a minimum number of spaces to ensure proper column alignment.

We can specify the minimum width after the percent sign for a formatting verb. If the argument matching that verb is shorter than the minimum width, it'll be padded with spaces until the required width is reached:

```
	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("--------------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)
```

```
     Product | Cost in Cents
--------------------------------
      Stamps | 50
 Paper Clips |  5
        Tape | 99
```

As the output shows us, we'll get padding wherever necessary according to our formatting rules!

## Formatting fractional number widths

Now, let's look at doing what we achieved with floats; using value widths to specify the precision for floating-point numbers:

```
%5.3f
```

Here's a breakdown of what the above says:

- % - Start of the formatting verb
- 5\. - Minimum width of the entire number
- 3 - Width **after** decimal point
- f - Formatting verb type

The minimum width of the entire number includes decimal places and the decimal point. If it's included, shorter numberse will be padded with spaces at the start until this width is reached. If it's omitted, no spaces will ever be added.

The width after the decimal point is the ubmer of decimal places to show.

Pretty useful stuff! Let's use our original program to use it:

```
func main() {
    width = 4.2
    height = 3.0
    area = width * height
    fmt.Printf("%.2f liters needed\n", area/10.0)

    width = 5.2
    height = 3.5
    area = width * height
    fmt.Printf("%.2f liters needed\n", area/10.0)
}
```

```
1.26 liters needed
1.82 liters needed
```

We successfully fixed our output to look reasonable. However, this is still pretty inefficient, so in the next file we'll look at more directly linked function stuff.
