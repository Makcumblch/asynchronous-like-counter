import { observer } from "mobx-react-lite";
import styles from "./App.module.css";
import { useMemo } from "react";
import { LikeCounterStorage } from "./storages/likeCounter.storage";

const App = observer(() => {
  
  const likeCounter = useMemo(() => {
    return new LikeCounterStorage();
  }, []);

  return (
    <>
      <h1>Ассинхронный счётчик лайков</h1>
      <div className={styles.card}>
        <p className={styles.count}>Число лайков: {likeCounter.count}</p>
        <button onClick={() => likeCounter.Increment()}>👍</button>
        <p>Нажмите на 👍 чтобы увеличить число лайков</p>
      </div>
    </>
  );
});

export default App;
