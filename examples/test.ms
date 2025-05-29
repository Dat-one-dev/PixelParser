// This is a single-line comment
/* This is a
   multi-line comment */

// Declare a color in RGBA format
color1 = [255, 0, 0, 255]      // Pure red with full opacity

// Declare a color in HEX format
color2 = #00FF00               // Pure green

// Convert HEX to RGBA
converted = rbga(color2)       // Now converted is in RGBA

// Push RGBA color to pixel (10, 10) on display 1
display(1).pixel(10,10).push converted

// Print out the color at that pixel in both RGBA and HEX
print pixel(10,10).color

// Fill a vertical line of pixels from (5,5) to (5,10) with red
display(1).pixel(5,5:5,10).push color1

// Print hello world for fun
print "Hello, Atomicity!"
