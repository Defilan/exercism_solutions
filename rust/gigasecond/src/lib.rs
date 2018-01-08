extern crate chrono;
use chrono::*;

// Returns a Utc DateTime one billion seconds after start.
pub fn after(start: DateTime<Utc>) -> DateTime<Utc> {
    let giganum = 10i64.pow(9);
    start + Duration::seconds(giganum)
}
