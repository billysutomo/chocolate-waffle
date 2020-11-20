import React from "react";
import styled from "styled-components";

const BasicLayoutStyled = styled.div<BasicLayoutProps>`
    height: 100vh;
    background-color: ${(p) => p.backgroundColor};
`

export interface BasicLayoutProps {
    children: React.ReactNode;
    backgroundColor: string;
}

const BasicLayout: React.FC<BasicLayoutProps> = ({
    children, 
    backgroundColor
}) => {
    return <BasicLayoutStyled backgroundColor={backgroundColor}>{children}</BasicLayoutStyled>;
};

export default BasicLayout;