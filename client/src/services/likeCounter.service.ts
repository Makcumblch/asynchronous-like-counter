import { clientIncrement, clientRead } from "./axios";

class LikeCounterService {
  static Get = () => {
    return clientRead.get("/likes");
  };
  static Increment = () => {
    return clientIncrement.post("/like/increment");
  };
}

export default LikeCounterService;
