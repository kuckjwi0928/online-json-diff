import React from "react";
import ReactDiffViewer from 'react-diff-viewer'

function DiffViewer(props) {
  const { left, right } = props
  return (
    <div style={{marginBottom: '65px'}}>
      <ReactDiffViewer
        oldValue={left}
        newValue={right}
        splitView
        useDarkTheme
      />
    </div>
  )
}

export default DiffViewer
