extern crate redis;

mod data_structs;
mod handlers;
mod ws;

use futures::{SinkExt, TryFutureExt};
use tokio::net::{TcpListener, TcpStream};

use crate::data_structs::UserDataRoot;
use charrid::{CharrID, CharredProcess};
use futures_util::{future, StreamExt, TryStreamExt};
use redis::{Commands, RedisResult};
use warp::{Filter, Rejection};

type Result<T> = std::result::Result<T, Rejection>;

#[tokio::main]
async fn main() -> RedisResult<()> {
    let mut proc = CharredProcess::new(255);

    let ws_route = warp::path("ws")
        .and(warp::ws())
        .and_then(handlers::ws_handler);
    let routes = ws_route.with(warp::cors().allow_any_origin());
    warp::serve(routes).run(([127, 0, 0, 1], 8000)).await;
    // let client = redis::Client::open("redis://127.0.0.1/")?;
    // let mut con = client.get_connection()?;
    // println!("ID: {:X}", id.as_u64());
    // let data: UserDataRoot = UserDataRoot {
    //     name: "Endercass",
    //     discrim_id: "0001",
    //     nickname: "Endercass",
    //     hash_pass: "$2y$10$aCjuStsnwL0V0P/XpSFDUO6UgfWiw7JiVEJbh1Q/SGEXEW1fz8TXG",
    // };
    // let _: () = con.hset_multiple(format!("user:{:X}", id.as_u64()), &data.as_tuple())?;

    Ok(())
}
