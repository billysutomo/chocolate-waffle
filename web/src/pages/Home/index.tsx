import React from 'react';
import { Block } from '../../components/Block';
import { BlockWrapper } from '../../components/BlockWrapper';

export default () => {
  return (
    <div>
      <Block active={false}>+ Add Block</Block>
      <BlockWrapper>
        <Block active={true} size="medium">+ Add Messenger</Block>
        <Block active={true} size="medium">+ Add Messenger</Block>
      </BlockWrapper>
      <BlockWrapper>
        <Block active={true} size="small">+ Add </Block>
        <Block active={true} size="small">+ Add </Block>
        <Block active={true} size="small">+ Add </Block>
      </BlockWrapper>
    </div>
  )
}