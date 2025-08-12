pub type Approval{
  Yes
  No
  Maybe
}

pub type Cuisine {
  Korean
  Turkish
}

pub type Genre {
  Crime
  Horror
  Romance
  Thriller
}

pub type Activity{
  BoardGame
  Chill
  Movie(Genre)
  Restaurant(Cuisine)
  Walk(Int)
}

pub fn rate_activity(activity: Activity) -> Approval {
  case activity{
    BoardGame -> No
    Chill -> No
    Movie(Romance) -> Yes
    Restaurant(Korean) -> Yes
    Restaurant(Turkish) -> Maybe
    Walk(distance) ->
      case distance > 11 {
        True -> Yes
        False -> 
          case distance > 6 {
            True -> Maybe
            False -> No
          }
      }
    _ -> No
  }
}
