import React from "react";

function ViewBadge(props) {
  const { items, deleteItem } = props
  return (
    <div className="mt-5">
      {
        Object.keys(items).map((key, index) =>
          <button className={`badge badge-outline ${index !== 0 ? 'ml-1.5' : ''}`} key={key} onClick={() => deleteItem(key)}>
            {`${key}=${items[key]}`}
          </button>
        )
      }
    </div>
  )
}

export default ViewBadge
