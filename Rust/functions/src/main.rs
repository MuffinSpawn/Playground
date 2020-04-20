fn main() {
    another_function(5, 6);

    let x = {
        five()
    };

    println!("The value of x is: {}", x);
}

fn another_function(x: i32, y: i32) {
    println!("The value of x is {}.", x);
    println!("The value of y is {}.", y);

    println!("Range: {:?}", (4..1))
}

fn five() -> i32 {
    5
}
