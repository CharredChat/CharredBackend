use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct UserDataRoot {
    pub name: &'static str,
    pub discrim_id: &'static str,
    pub nickname: &'static str,
    pub hash_pass: &'static str,
}

impl UserDataRoot {
    pub fn as_tuple(&self) -> [(&str, &str); 4] {
        [
            ("name", self.name),
            ("discrim_id", self.discrim_id),
            ("nickname", self.nickname),
            ("hash_pass", self.hash_pass),
        ]
    }
}
