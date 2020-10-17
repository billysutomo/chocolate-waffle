import React from "react";
import { Block } from "../../components/Block";
import { BlockWrapper } from "../../components/BlockWrapper";

const Component: React.FC = () => {
  return (
    <div>
      <Block active={false}>+ Add Block</Block>
      <BlockWrapper>
        <Block active size="medium">
          + Add Messenger
        </Block>
        <Block active size="medium">
          + Add Messenger
        </Block>
      </BlockWrapper>
      <BlockWrapper>
        <Block active size="small">
          + Add
        </Block>
        <Block active size="small">
          + Add
        </Block>
        <Block active size="small">
          + Add
        </Block>
      </BlockWrapper>
    </div>
  );
};

export default Component;
