import React from 'react';
// import CameraButton from '../../components/CameraButton';
import { BlockButton } from '../../components/BlockButton';

export default () => {
  return (
    <div>
      <BlockButton active={true}>+ Add Messenger</BlockButton>
      <BlockButton active={false}>+ Add Block</BlockButton>
    </div>
  )
}