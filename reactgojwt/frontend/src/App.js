import "./App.css";
import NavbarTrue from "./components/NavbarTrue";
import FormLogin from "./pages/FormLogin";

import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import FormRegister from "./pages/FormRegister";
import React, { useEffect, useState } from "react";

function App() {
  const [name, setName] = useState("");

  useEffect(() => {
    (async () => {
      const response = await fetch("http://localhost:8000/api/user", {
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
      });
      const data = await response.json();
      setName(data.name);
    })();
  }, []);


  

  return (
    <div className="App">
      <Router>
        <NavbarTrue />
        <h1> {name ? "Hi " + name : "Youre not login"}</h1>

        <Switch>
          <Route exact path="/">
            {/* <Form /> */}
          </Route>
          <Route path="/login">
            <FormLogin />
          </Route>
          <Route path="/register">
            <FormRegister />
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
