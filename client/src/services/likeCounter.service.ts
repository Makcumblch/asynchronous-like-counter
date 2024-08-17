import client from "./axios";

class LikeCounterService {
  static Get = () => {
    return client.get("/likes");
  };
  static Increment = () => {
    return client.post("/like/increment");
  };
}

export default LikeCounterService;
