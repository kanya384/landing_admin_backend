import "./spinner.css"

const Spinner = () => {
  return(
    <main className="main">
      <div className="loading">
        <div className="spinner-border text-primary" role="status"><span className="visually-hidden">Loading...</span></div>
      </div>
    </main>
  )
}

export default Spinner