use std::env;
use rusoto_core::{HttpClient};
use rusoto_core::region::Region;
use rusoto_s3::{S3, S3Client, ListObjectsV2Request};
use rusoto_credential::{StaticProvider};

pub fn main() {
    let key_id = env::var("AWS_ACCESS_KEY_ID").unwrap();
    let secret = env::var("AWS_SECRET_ACCESS_KEY").unwrap();
    let region = env::var("AWS_DEFAULT_REGION").unwrap();
    let bucket = env::var("S3_BUCKET_NAME").unwrap();

    let custom_region = Region::Custom {
        name: region.to_owned(),
        endpoint: format!("s3.{}.amazonaws.com", region.to_owned()),
    };
    let provider = StaticProvider::new_minimal(key_id, secret);
    let client = S3Client::new_with(
        HttpClient::new().unwrap(),
        provider,
        custom_region,
    );
    let list_obj_req = ListObjectsV2Request {
        bucket: bucket,
        ..Default::default()
    };
    match client.list_objects_v2(list_obj_req).sync() {
        Ok(val) => println!("{:#?}", val),
        Err(err) => eprintln!("{:#?}", err),
    }
}
