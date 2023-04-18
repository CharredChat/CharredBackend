use std::collections::HashMap;

pub struct user_data_root {
    pub name: &'static str,
    pub discrim_id: &'static str,
    pub nickname: &'static str,
    pub hash_pass: &'static str,
}

impl user_data_root {
    pub fn as_tuple(&self) -> [(&str, &str); 4] {
        [
            ("name", self.name),
            ("discrim_id", self.discrim_id),
            ("nickname", self.nickname),
            ("hash_pass", self.hash_pass),
        ]
    }
}
