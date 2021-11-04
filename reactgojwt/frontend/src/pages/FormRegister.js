import React, { useState } from "react";
import { Redirect } from "react-router-dom";

import { Form, FormGroup, Label, Button, Input } from "reactstrap";

const FormRegister = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [redirect, setredirect] = useState(false);

  const submit = async (e) => {
    e.preventDefault();
    console.log(name, email, password);
    await fetch("http://localhost:8000/api/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name,
        email,
        password,
      }),
    });
    setredirect(true);
  };

  if (redirect) {
    return <Redirect to="/login" />;
  }

  //   <Redirect to="/login"/>

  return (
    <div>
      <Form inline className="" onSubmit={submit}>
        <FormGroup floating>
          <Input
            name="name"
            placeholder="Name"
            type="text"
            required
            onChange={(e) => setName(e.target.value)}
          />
          <Label>Name</Label>
        </FormGroup>
        <FormGroup floating>
          <Input
            id="exampleEmail"
            name="email"
            placeholder="Email"
            type="email"
            required
            onChange={(e) => setEmail(e.target.value)}
          />
          <Label for="exampleEmail">Email</Label>
        </FormGroup>
        <FormGroup floating>
          <Input
            id="examplePassword"
            name="password"
            placeholder="Password"
            type="password"
            required
            onChange={(e) => setPassword(e.target.value)}
          />
          <Label for="examplePassword">Password</Label>
        </FormGroup>
        <Button>Submit</Button>
      </Form>
    </div>
  );
};

export default FormRegister;
