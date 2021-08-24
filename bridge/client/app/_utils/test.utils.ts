export class TestUtils {
  public static createNewDropEventWithFiles(files: File[]): DragEvent {
    const dataTransfer: DataTransfer = MockDataTransfer(files);
    const event = new DragEvent('drop');
    Object.defineProperty(event.constructor.prototype, 'dataTransfer', {
      value: dataTransfer
    });
    Object.defineProperty(event.constructor.prototype, 'preventDefault', {
      value: () => { return; }
    });
    Object.defineProperty(event.constructor.prototype, 'stopPropagation', {
      value: () => { return; }
    });
    return event;
  }

  public static mockWindowMatchMedia() {
    Object.defineProperty(window, 'matchMedia', {
      writable: true,
      value: jest.fn().mockImplementation(query => ({
        matches: false,
        media: query,
        onchange: null,
        addListener: jest.fn(), // Deprecated
        removeListener: jest.fn(), // Deprecated
        addEventListener: jest.fn(),
        removeEventListener: jest.fn(),
        dispatchEvent: jest.fn(),
      })),
    });
  }
}

function MockDataTransfer(files: File[]): DataTransfer {
  return {
    // @ts-ignore
    dropEffect: 'all',
    effectAllowed: 'all',
    // @ts-ignore
    items: [],
    types: ['Files'],
    // @ts-ignore
    getData() {
      return files;
    },
    // @ts-ignore
    files: [...files]
  };
}