import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";

const BlockWrapperStyled = styled.div`
  display: flex;
  justify-content: space-between;
  width: 370px;
`;

export interface BlockWrapperProps {
  children: React.ReactNode;
}

export const BlockWrapper: React.FC<BlockWrapperProps> = ({ children }) => {
  return <BlockWrapperStyled>{children}</BlockWrapperStyled>;
};

BlockWrapper.propTypes = {
  children: PropTypes.node.isRequired,
};
