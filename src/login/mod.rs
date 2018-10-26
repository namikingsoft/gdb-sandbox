use std::process;
use crypto::sha2::Sha256;
use crypto::hmac::Hmac;
use crypto::mac::Mac;
use rpassword;

fn password_hash(password: String) -> String {
    let key = env!("CIPHER_KEY_PHRASE").to_string();
    let mut hmac = Hmac::new(Sha256::new(), key.as_bytes());
    hmac.input(password.as_bytes());
    let result = hmac.result();
    format!("{:x?}", result.code()).replace(", ", "").replace("[", "").replace("]","")
}

pub fn main() {
    let password = rpassword::prompt_password_stdout("Password: ").unwrap();
    let hash = password_hash(password.to_owned());
    println!("");
    println!("Input: {}", password);
    println!("SHA256: {}", hash);
    println!("");
    if hash != "117f79884be68260f69f288d9e1ac75338dd859a3c67d0da926817156c5a6442" {
        eprintln!("Login Failed\n");
        process::exit(0x01);
    }
    println!("Login Succeeed!\n");
}
