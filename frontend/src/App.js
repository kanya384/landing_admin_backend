import { SectionsContext } from "./context/sectionsContext"
import { useBlocks } from "./hooks/sections.hook"
import { BrowserRouter as Router, Route, Routes } from "react-router-dom"
import { useSections } from "./hooks/sections.hook";
import Main from "./pages/main";
import { Plans } from "./pages/plans";
import { PlanDetail } from "./pages/plan-detail";
import { Promo } from "./pages/promo";
import { About } from "./pages/about";

function App() {
  const { blocks, setBlocks, menuClick, setMenuClick } = useSections()
  return (
    <div className="App container_main">
       <div className="over">
        <SectionsContext.Provider value={{ blocks, setBlocks, menuClick, setMenuClick }}>
          <Router>
            <Routes>
             <Route path="/about" element={<About />} />
              <Route path="/promo" element={<Promo />} />
              <Route path="/plans/:id" element={<PlanDetail />} />
              <Route path="/plans" element={<Plans />} />
              <Route path="*" element={<Main />} />
            </Routes>
          </Router>
        </SectionsContext.Provider>
      </div>
    </div>
  );
}

export default App;