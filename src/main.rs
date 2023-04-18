extern crate redis;

mod data_structs;

use charrid::{CharrID, CharredProcess};
use redis::{Commands, RedisResult};

use crate::data_structs::user_data_root;

fn main() -> RedisResult<()> {
    let client = redis::Client::open("redis://127.0.0.1/")?;
    let mut proc = CharredProcess::new(255);
    let mut con = client.get_connection()?;
    let id = proc.generate();
    println!("ID: {:X}", id.as_u64());
    let data: user_data_root = user_data_root {
        name: "Endercass",
        discrim_id: "0001",
        nickname: "Endercass",
        hash_pass: "$2y$10$aCjuStsnwL0V0P/XpSFDUO6UgfWiw7JiVEJbh1Q/SGEXEW1fz8TXG",
    };
    let _: () = con.hset_multiple(format!("user:{:X}", id.as_u64()), &data.as_tuple())?;
    Ok(())
}
