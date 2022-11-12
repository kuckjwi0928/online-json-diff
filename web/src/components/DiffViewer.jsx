import React from "react";
import ReactDiffViewer from 'react-diff-viewer'

function DiffViewer(props) {
  const { left, right } = props
  return (
    <ReactDiffViewer
      oldValue={left}
      newValue={right}
      splitView
      useDarkTheme
    />
  )
}

export default DiffViewer
