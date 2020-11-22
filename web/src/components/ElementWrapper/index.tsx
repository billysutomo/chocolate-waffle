import React from "react";
import PropTypes from "prop-types";
import styled from "styled-components";

const ElementWrapperStyled = styled.div`
  display: block;
  max-width: 386px;
`;

export interface ElementWrapperProps {
  children: React.ReactNode;
}

export const ElementWrapper: React.FC<ElementWrapperProps> = ({ children }) => {
  return <ElementWrapperStyled>{children}</ElementWrapperStyled>;
};

ElementWrapper.propTypes = {
  children: PropTypes.node.isRequired,
};
