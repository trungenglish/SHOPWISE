import React from 'react';
import type { DynamicUIDocument, ComponentNode } from '@shopwise/ui-protocol/types';
import { validateDynamicUIDocument } from './validation';
// Import renderers (MVP placeholder approach)
import * as Commerce from './commerce';
import * as AI from './ai';
import * as Interaction from './interaction';
import * as ErrorUI from './error';

const componentRegistry: Record<string, React.FC<any>> = {
  ...Commerce,
  ...AI,
  ...Interaction,
  ...ErrorUI,
};

interface ProtocolRendererProps {
  document: any; // Raw JSON payload
}

export const ProtocolRenderer: React.FC<ProtocolRendererProps> = ({ document }) => {
  // Validate schema (for now using base schema or any mock schema)
  // In a real app we'd load the full JSON schema
  const { valid, errors } = validateDynamicUIDocument(document, {});

  if (!valid) {
    console.error("Invalid Dynamic UI Document", errors);
    return <ErrorUI.ErrorState title="UI Error" message="Failed to render content due to invalid data structure." severity="critical" />;
  }

  const doc = document as DynamicUIDocument;

  // Fallback UI handler for unsupported schema versions
  if (doc.schemaVersion !== "1.0") {
    return (
      <ErrorUI.ErrorState
        title="Version Mismatch"
        message={`Unsupported UI Schema Version: ${doc.schemaVersion}. Please update your client.`}
        severity="warning"
      />
    );
  }

  const renderNode = (node: ComponentNode): React.ReactNode => {
    const Component = componentRegistry[node.type];

    if (!Component) {
      console.warn(`Unsupported component type: ${node.type}`);
      return null;
    }

    return (
      <Component key={node.id} id={node.id} {...node.props} {...node.state}>
        {node.children?.map(child => renderNode(child))}
      </Component>
    );
  };

  return (
    <div className="dynamic-ui-root" data-document-id={doc.documentId}>
      {renderNode(doc.root)}
    </div>
  );
};
