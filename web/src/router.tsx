import React from "react";
import { Route, BrowserRouter as Router, Switch } from "react-router-dom";
import Home from "./pages/Home";
import Setting from "./pages/Setting";

const Component: React.FC = () => {
  return (
    <Router>
      <Switch>
        <Route exact path="/" component={Home} />
        <Route exact path="/setting" component={Setting} />
      </Switch>
    </Router>
  );
};

export default Component;
