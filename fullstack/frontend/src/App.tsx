import "./App.css";
import Header from "./components/Header";
import Todolist from "./components/Todolist";

function App() {
  return (
    <div className="App">
      <Header />
      <Todolist
        todos={[
          {
            title: "Cooking",
            desctopstion: "cooking with mama",
            inCompleted: false,
          },
          {
            title: "Solve Bug",
            desctopstion: "Complete Task",
            inCompleted: true,
          },

        ]}
      />
    </div>
  );
}

export default App;
