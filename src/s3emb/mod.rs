use rusoto_core::{HttpClient};
use rusoto_core::region::Region;
use rusoto_s3::{S3, S3Client, ListObjectsV2Request};
use rusoto_credential::{StaticProvider};

pub fn main() {
    let key_id = env!("AWS_ACCESS_KEY_ID").to_string();
    let secret = env!("AWS_SECRET_ACCESS_KEY").to_string();
    let region = env!("AWS_DEFAULT_REGION").to_string();
    let bucket = env!("S3_BUCKET_NAME").to_string();

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
