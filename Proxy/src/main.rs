trait Internet{
    fn connect_to(&self, address: String);
}

struct RealInternet {}

impl Internet for RealInternet {
    fn connect_to(&self, address: String) {
        println!("Connected to address: {}", address)
    }
}

struct ProxyInternet {
    real: RealInternet,
    banned_addr: Vec<String>,
}

impl Internet for ProxyInternet {
    fn connect_to(&self, address: String) {
        if self.is_banned(&address){
            println!("Address {} is banned", &address);
            return
        }

        self.real.connect_to(address)
    }
}

impl ProxyInternet {
    fn ban_addr(&mut self, address: String) {
        self.banned_addr.push(address)
    }

    fn is_banned(&self, address: &String) -> bool {
        if self.banned_addr.contains(&address) {
            return true
        }

        false
    }
}

fn main() {
    let real_i: RealInternet = RealInternet{};
    real_i.connect_to(String::from("abc.com"));

    let mut proxy_i: ProxyInternet = ProxyInternet{
        real: real_i,
        banned_addr: vec![],
    };
    proxy_i.ban_addr(String::from("abc.com"));
    proxy_i.connect_to(String::from("abc.com"));
    proxy_i.connect_to(String::from("cba.com"));
}
