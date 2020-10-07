import React from 'react';
import styled from 'styled-components';

export interface BlockButtonProps {
  label: string;
}

export const BlockButton: React.FC<BlockButtonProps> = ({
  label,
  ...props
}) => {
  return (
    <BlockButtonStyled {...props}>{label}</BlockButtonStyled>
  )
}

const BlockButtonStyled = styled.button`
  border: 2px dashed #607D8B;
  border-radius: 10px;
  padding: 15px 20px 15px 20px;
  width: 300px;
  text-align: center;
  :hover{
    cursor: pointer;
  }
  :active{
    background: #2ab8b96b;
  }
  :focus{
    outline: 0;
  }
`