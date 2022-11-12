import React from "react";

function BottomNav() {
  return (
    <div className="btm-nav">
      <button onClick={() => window.scrollTo(0, 0)}>Top</button>
    </div>
  )
}

export default BottomNav
