use charrid::CharredProcess;
use futures::{FutureExt, SinkExt, StreamExt};
use warp::ws::{Message, WebSocket};

use crate::data_structs::UserDataRoot;

pub async fn client_connection(ws: WebSocket, proc: &mut CharredProcess) {
    let (mut tx, mut rx) = ws.split();
    while let Some(result) = rx.next().await {
        tx.send(Message::text(format!(
            "Echo: {}\nUser Example: {}",
            result.expect("msg").to_str().expect("msg"),
            proc.generate().as_u64()
        )))
        .await
        .expect("msg");
    }
}
