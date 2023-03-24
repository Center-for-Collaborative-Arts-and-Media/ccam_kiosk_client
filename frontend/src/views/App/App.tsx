import type { Component } from "solid-js";
import styles from "./App.module.css";
import Slideshow from "../../components/slideshow/Slideshow";

const App: Component = () => {
  return (
    <div class={styles.App}>
      <Slideshow />
    </div>
  );
};

export default App;
