use core::fmt;

#[derive(Debug)]
pub struct Clock {
    hour: i32,
    mins: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let mut total = (hours * 60) + minutes;
        total = total % (24 * 60);
        if total < 0 {
            let hour = ((24 * 60) + total) / 60;
            let mut mins = minutes % 60;
            if mins < 0 {
                mins = 60 + mins;
            }
            Clock {
                hour: hour,
                mins: mins,
            }
        } else {
            let hour = total / 60;
            let mins = total % 60;
            Clock {
                hour: hour,
                mins: mins,
            }
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        let mut temp = Clock {
            hour: self.hour,
            mins: self.mins,
        };

        let mut total_mins = (temp.hour * 60) + temp.mins;

        total_mins = total_mins + minutes;
        total_mins = total_mins % (24 * 60);
        if total_mins < 0 {
            total_mins = (24 * 60) + total_mins
        }
        temp.hour = total_mins / 60;
        temp.mins = total_mins - (temp.hour * 60);
        temp
    }
}

impl PartialEq for Clock {
    fn eq(&self, other: &Self) -> bool {
        self.hour == other.hour && self.mins == other.mins
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:02}:{:02}", self.hour, self.mins)
    }
}