export const Parallax = ({image}) => {
  return ( <section className="js-parrallax parrallax-scene1">
              <div className="fake-img fake-img1"><img src={image} alt="" /></div>
              <div id="scrollEl1" className="parralax-img1 img-hidden"><img src={image} alt="" /></div>
            </section>
  )
}