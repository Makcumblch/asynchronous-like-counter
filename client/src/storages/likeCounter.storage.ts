import { makeAutoObservable, runInAction } from "mobx";
import LikeCounterService from "../services/likeCounter.service";

const syncTimeout = 5000;

export class LikeCounterStorage {
  count: number = 0;
  constructor() {
    makeAutoObservable(this);

    this.getCount();

    setInterval(() => {
      this.getCount();
    }, syncTimeout);
  }

  getCount = () => {
    LikeCounterService.Get().then((value) => {
      if (!value.data) return;
      runInAction(() => {
        this.count = value.data;
      });
    });
  };

  Increment = () => {
    this.count += 1;
    LikeCounterService.Increment().catch(() => {
      this.decrement();
    });
  };

  decrement = () => {
    this.count -= 1;
  };
}
