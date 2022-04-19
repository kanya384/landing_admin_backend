import { useTypedSelector } from "../../hooks/use-typed-selector"
import { Authentication } from "../../pages/authentication"
import { BrowserRouter, Route, Routes } from "react-router-dom"
import Spinner from "../../components/spinner/spinner"
import { useEffect } from "react"
import { useActions } from "../../hooks/use-actions"

const Router: React.FC = () => {
  const auths = useTypedSelector(({ auths }) => {
    return auths
  })
  const { checkAuth } = useActions();

  useEffect(()=>{
    checkAuth()
  }, [checkAuth])

  if (auths?.loading === true) {
    return <Spinner />
  } else if (auths?.auth === true) {
    return (
      <BrowserRouter>
        <Routes>
          <Route path="*" element={<div>Hello world</div>} />
        </Routes>
      </BrowserRouter>
    )
  } else {
    return (
      <BrowserRouter>
        <Routes>
          <Route path="*" element={<Authentication />} />
        </Routes>
      </BrowserRouter>
    )
  }
}

export default Router