/* @refresh reload */
import { render } from "solid-js/web";
import { Router, Route, Routes } from "@solidjs/router";

import "./styles/reset.css";
import "./styles/index.css";
import App from "./views/App/App";
import Slides from "./views/Slides/Slides";

render(
  () => (
    <Router>
      <Routes>
        <Route path="/" component={App} />
        <Route path="/slides" component={Slides} />
        <Route path="*" component={() => <h1>404</h1>} />
      </Routes>
    </Router>
  ),
  document.getElementById("root") as HTMLElement
);
